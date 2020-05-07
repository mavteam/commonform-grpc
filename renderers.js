var docx = require('commonform-docx')
var html = require('commonform-html')
var mark = require('commonform-commonmark')

module.exports = {
  html: async (form, blanks, options) => {
    return html(form, blanks, options)
  },

  html5: async (form, blanks, options) => {
    options = options || {}
    options.html5 = true
    return html(form, blanks, options)
  },

  docx: async (form, blanks, options) => {
    return docx(form, blanks, options).generateAsync({ type: 'nodebuffer' })
  },

  markdown: async (form, blanks, options) => {
    return mark.stringify(form, blanks, options)
  }
}
