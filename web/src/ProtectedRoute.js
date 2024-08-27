import { Navigate } from 'react-router-dom'
import { getToken } from 'data/auth'

// eslint-disable-next-line react/prop-types
const ProtectedRoute = ({ children }) => {
  if (!getToken()) {
    return <Navigate to="/login" replace />
  }
  return children
}

export default ProtectedRoute
