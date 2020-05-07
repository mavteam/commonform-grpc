const parseForm = require('./parser')
const renderers = require('./renderers')
const numberings = require('./numberings')

const docxMime = 'application/vnd.openxmlformats-officedocument.wordprocessingml.document'
const htmlMime = 'text/html'
const markdownMime = 'text/markdown'

const Extract = async (call) => {
  // Parse form
  const parsed = await parseForm(call.request.data.toString('utf8'), call.request.meta.format)
    .catch((error) => {
      throw error
    })

  // Process blanks
  const result = { blanks: [] }
  for (var b of parsed.directions) {
    result.blanks.push(b.label)
  };

  return result
}

const Assemble = async (call) => {
  // Parse form
  const parsed = await parseForm(call.request.document.data.toString('utf8'), call.request.document.meta.format)
    .catch((error) => {
      throw error
    })

  // Process options
  const passthroughOptionKeys = [
    'title', 'hash', 'indentMargins', 'markFilled'
  ]
  const options = {}
  passthroughOptionKeys.forEach(function (key) {
    if (call.request[key]) {
      options[key] = call.request[key]
    }
  })
  options.blanks = call.request.unfilledBlanks
  options.centerTitle = !call.request.leftAlignTitle

  // Process styles
  // See commonform.proto for how to send styles selection in.
  // Node panics if this JSON string is an empty buffer.
  if (call.request.styles.length !== 0) {
    options.styles = JSON.parse(call.request.styles)
  }

  // Process numbering
  // See commonform.proto for how to send numbering selection in.
  options.numbering = numberings[call.request.numbering]

  // Process blanks
  // See commonform.proto for how to send blanks in.
  let blanks = {}
  if (call.request.blanks.length !== 0) {
    for (const b of call.request.blanks) {
      blanks[b.name] = b.value
    }
  };
  if (parsed.directions && !Array.isArray(blanks)) {
    blanks = Object.keys(blanks).reduce(function (output, key) {
      parsed.directions
        .filter(function (direction) {
          return direction.label === key
        })
        .forEach(function (direction) {
          output.push({
            label: direction.label,
            blank: direction.blank,
            value: blanks[key]
          })
        })
      return output
    }, [])
  }

  // Process signatures
  // See commonform.proto for how to send signatures in.
  if (call.request.format === 'docx' && (call.request.useExternalSignatures)) {
    options.after = require('./lib/ooxml-signature-pages-external')(call.request.externalSignatureCount)
  } else if (call.request.format === 'docx' && call.request.signatures) {
    options.after = require('ooxml-signature-pages')(call.request.signatures)
  }

  // Render
  const renderer = renderers[call.request.format]
  const rendered = await renderer(parsed.form, blanks, options)
    .catch((error) => {
      throw error
    })
  const result = {
    meta: {
      name: call.request.document.meta.name
    },
    data: Buffer.from(rendered)
  }

  // Set mimetype
  switch (call.request.format) {
    case 'docx':
      result.meta.mime = docxMime
      break
    case 'html':
    case 'html5':
      result.meta.mime = htmlMime
      break
    case 'markdown':
      result.meta.mime = markdownMime
      break
  }

  return result
}

module.exports = {
  Extract,
  Assemble
}
