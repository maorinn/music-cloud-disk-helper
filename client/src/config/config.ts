const IS_DEBUG = process.env.NODE_ENV === "development"
export const config = {
    SERVER_HOME:IS_DEBUG?"127.0.0.1:22333":`${window.location.hostname}:22333`
}