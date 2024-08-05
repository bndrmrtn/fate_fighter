/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./ui/**/*.{html,js,templ}"
  ],
  theme: {
    extend: {
      colors: {
        dark: '#1d1d27',
      }
    },
  },
  plugins: [
    require('@tailwindcss/forms')
  ],
  darkMode: ['selector', '.dark']
}

