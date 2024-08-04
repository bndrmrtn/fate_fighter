/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./ui/**/*.{html,js,templ}"
  ],
  theme: {
    extend: {},
  },
  plugins: [],
  darkMode: ['selector', '[data-mode="dark"]']
}

