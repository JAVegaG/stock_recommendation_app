declare global {
    interface Window {
        env?: Record<string, string>;
    }
}

export const getEnv = (env?: string) => {
    if (!window) return undefined
    if (!window.env) return undefined

    if (!env) return window.env

    return window.env[env]
}