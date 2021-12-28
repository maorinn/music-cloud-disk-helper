const IS_DEBUG = process.env.NODE_ENV === "development"
export const config = {
    SERVER_HOME:IS_DEBUG?"127.0.0.1:8000":`${window.location.hostname}:8000`
}