import './Header.scss'
import { Link, useLocation } from 'react-router-dom'
import { MENU_ITEMS } from 'utils/contants'
import classNames from 'classnames'
import Notification from './Notification'
import UserInfo from './UserInfo'
import Images from 'images'
import usePremiseParam from 'hooks/usePremiseParam'

const Header = () => {
  const { pathname } = useLocation()
  const { withPremiseParam } = usePremiseParam()

  return (
    <div className="header">
      <div className="header__logo">
        <img src={Images.appIcon} alt="logo" />
      </div>
      <div className="header__title">
        <Link to={withPremiseParam('/')}>Security Command Center</Link>
      </div>
      <div className="header__menu">
        {MENU_ITEMS.map(({ name, path, isMultiPath }, index) => (
          <Link
            key={index}
            to={withPremiseParam(path)}
            className={classNames({
              'is-active': isMultiPath ? pathname.includes(path) : pathname === path
            })}>
            {name}
          </Link>
        ))}
      </div>

      <div className="header__menu--right">
        <Notification />
        <UserInfo />
      </div>
    </div>
  )
}

export default Header
