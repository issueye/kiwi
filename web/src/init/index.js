import { installElementIcons, installComponents, installColipboard } from './install_components'

export function install(Vue) {
    // 安装图标
    installElementIcons(Vue)

    // 安装组件
    installComponents(Vue)

    // 安装剪贴板
    installColipboard(Vue)
}