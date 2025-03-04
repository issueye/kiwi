function genSimilarColorsName(brandName) {
  return {
    DEFAULT: `var(--el-color-${brandName})`,
    500: `var(--el-color-${brandName})`,
    400: `var(--el-color-${brandName}-dark-2)`,
    600: `var(--el-color-${brandName}-light-3)`,
    700: `var(--el-color-${brandName}-light-5)`,
    800: `var(--el-color-${brandName}-light-7)`,
    900: `var(--el-color-${brandName}-light-9)`,
  };
}

/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: "class",
  corePlugins: {
    preflight: false
  },
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        bg_color: "var(--el-bg-color)",
        primary: genSimilarColorsName('primary'),
        info: genSimilarColorsName('info'),
        success: genSimilarColorsName('success'),
        warning: genSimilarColorsName('warning'),
        danger: genSimilarColorsName('danger'),
        text_color_primary: "var(--el-text-color-primary)",
        text_color_regular: "var(--el-text-color-regular)",
        text_color_disabled: "var(--el-text-color-disabled)",
      }
    }
  }
};