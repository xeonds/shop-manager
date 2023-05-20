export default function isLogin() {
    if (localStorage.getItem('token') == null) return false
    if (localStorage.getItem('token') == '') return false
    // parse token
    const token = JSON.parse(localStorage.getItem('token') as string)
    // check expired
    if (Date.now() > token.expire) return false
    return true
}