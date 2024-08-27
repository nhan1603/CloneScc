import { getToken } from 'data/auth'
import './Login.scss'
import LoginForm from './LoginForm'
import Images from 'images'
import { Navigate } from 'react-router-dom'

const Login = () => {
  // redirect to dashboard if user already login
  if (getToken()) return <Navigate to="/" />

  return (
    <div className="container">
      <div className="logo">
        <img src={Images.loginLogo} alt="sp-logo"></img>
      </div>
      <div className="login-form">
        <img src={Images.loginImg} alt="form-logo"></img>
        <LoginForm />
      </div>
    </div>
  )
}

export default Login
