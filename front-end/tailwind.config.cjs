/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './index.html',
    './src/**/*.{vue,ts,html}'
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('tailwindcss-debug-screens'),
    require('tailwind-scrollbar')
  ],
}
