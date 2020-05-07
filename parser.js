var commonmark = require('commonform-commonmark')
var parseMarkup = require('commonform-markup-parse')

module.exports = async (formData, format) => {
  /* istanbul ignore else */
  if (format === 'json') {
    return { form: JSON.parse(formData), directions: [], frontmatter: {} }
  } else if (format === 'markup') {
    return parseMarkup(formData)
  } else if (format === 'commonmark') {
    return commonmark.parse(formData)
  }
}
