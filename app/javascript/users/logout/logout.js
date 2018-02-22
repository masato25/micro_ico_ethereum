import Cookies from 'js-cookie';
const $ = require("jQuery")

$().ready(() => {
  Cookies.remove('auth_session')
  setTimeout(() => {
    window.location.href = "/login"
  }, 500)
  console.log("delete auth_session cookie")
})
