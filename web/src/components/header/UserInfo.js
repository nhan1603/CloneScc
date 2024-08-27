import { useRef, useState } from 'react'
import './UserInfo.scss'
import useClickOutside from 'hooks/useClickOutside'
import { getCurrentUser, removeToken } from 'data/auth'
import { useNavigate } from 'react-router-dom'

const UserInfo = () => {
  const [isOpen, setIsOpen] = useState(false)
  const user = getCurrentUser()
  const navigate = useNavigate()
  const ref = useRef()

  useClickOutside({
    isOpen,
    ref,
    handleClose: () => {
      setIsOpen(false)
    }
  })

  return (
    <div
      id="header-user-info"
      className="user-info"
      data-testid="user-info"
      ref={ref}
      onClick={() => {
        if (!isOpen) {
          setIsOpen(true)
        }
      }}>
      <span className="user">{user?.email?.at(0)}</span>
      {isOpen && (
        <div
          className="logout"
          data-testid="logout"
          onClick={() => {
            removeToken()
            navigate('/login')
          }}>
          Logout
        </div>
      )}
    </div>
  )
}

export default UserInfo
