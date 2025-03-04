import Color from 'color'
import { useUserStore } from '~/store'

// 主题管理类
export class ThemeManager {
  // 生成主题的衍生色
  static generateThemeColors(primaryColor) {
    const primary = Color(primaryColor)
    const hue = primary.hue()
    const saturation = primary.saturationl()
    
    // 生成主色调的变体
    const colors = {
      'primary': primaryColor,
      'primary-dark-2': primary.darken(0.2).hex(),
      'primary-light-3': primary.lighten(0.3).hex(),
      'primary-light-5': primary.lighten(0.5).hex(),
      'primary-light-7': primary.lighten(0.7).hex(),
      'primary-light-8': primary.lighten(0.8).hex(),
      'primary-light-9': primary.lighten(0.9).hex(),

      // 成功色：偏绿色调
      'success': Color.hsl((hue + 100) % 360, saturation * 0.9, 45).hex(),
      'success-dark-2': Color.hsl((hue + 100) % 360, saturation * 0.8, 40).hex(),
      'success-light-3': Color.hsl((hue + 100) % 360, saturation * 0.8, 55).hex(),
      'success-light-5': Color.hsl((hue + 100) % 360, saturation * 0.7, 65).hex(),
      'success-light-7': Color.hsl((hue + 100) % 360, saturation * 0.5, 75).hex(),
      'success-light-8': Color.hsl((hue + 100) % 360, saturation * 0.4, 80).hex(),
      'success-light-9': Color.hsl((hue + 100) % 360, saturation * 0.3, 85).hex(),

      // 信息色：固定灰色调
      'info': '#909399',
      'info-dark-2': '#828489',
      'info-light-3': '#B1B3B8',
      'info-light-5': '#C8C9CC',
      'info-light-7': '#DEDFE0',
      'info-light-8': '#E9E9EB',
      'info-light-9': '#F4F4F5',

      // 警告色：偏橙色调
      'warning': Color.hsl((hue + 30) % 360, saturation * 0.95, 55).hex(),
      'warning-dark-2': Color.hsl((hue + 30) % 360, saturation * 0.9, 50).hex(),
      'warning-light-3': Color.hsl((hue + 30) % 360, saturation * 0.85, 65).hex(),
      'warning-light-5': Color.hsl((hue + 30) % 360, saturation * 0.75, 75).hex(),
      'warning-light-7': Color.hsl((hue + 30) % 360, saturation * 0.55, 85).hex(),
      'warning-light-8': Color.hsl((hue + 30) % 360, saturation * 0.45, 90).hex(),
      'warning-light-9': Color.hsl((hue + 30) % 360, saturation * 0.35, 95).hex(),

      // 危险色：偏红色调
      'danger': Color.hsl((hue - 30) % 360, saturation * 0.9, 60).hex(),
      'danger-dark-2': Color.hsl((hue - 30) % 360, saturation * 0.85, 55).hex(),
      'danger-light-3': Color.hsl((hue - 30) % 360, saturation * 0.8, 70).hex(),
      'danger-light-5': Color.hsl((hue - 30) % 360, saturation * 0.7, 80).hex(),
      'danger-light-7': Color.hsl((hue - 30) % 360, saturation * 0.5, 90).hex(),
      'danger-light-8': Color.hsl((hue - 30) % 360, saturation * 0.4, 92).hex(),
      'danger-light-9': Color.hsl((hue - 30) % 360, saturation * 0.3, 95).hex()
    }
    return colors
  }

  // 应用主题颜色
  static applyTheme(colors) {
    const root = document.documentElement
    Object.entries(colors).forEach(([key, value]) => {
      root.style.setProperty(`--el-color-${key}`, value)
      // 同步更新Tailwind颜色变量
      if (key === 'primary') {
        root.style.setProperty('--color-primary-500', value)
        root.style.setProperty('--tw-text-opacity', '1')
        root.style.setProperty('--tw-bg-opacity', '1')
      }
    })
  }

  // 设置主题颜色
  static setTheme(primaryColor) {
    const colors = this.generateThemeColors(primaryColor)
    this.applyTheme(colors)
    // 读取 pinia 的 themeColor
    useUserStore().themeColor = primaryColor

    // localStorage.setItem('theme-primary-color', primaryColor)
  }

  // 初始化主题
  static initTheme() {
    // const savedColor = localStorage.getItem('theme-primary-color')
    // 读取 pinia 的 themeColor
    const savedColor = useUserStore().themeColor
    if (savedColor) {
      this.setTheme(savedColor)
    }
  }
}