const IS_DEBUG = process.env.NODE_ENV
export const config = {
    SERVER_HOME:IS_DEBUG?"127.0.0.1:8000":"42.192.50.25:8000"
}