import Header from 'components/header'
import PropTypes from 'prop-types'

const CommonLayout = ({ children }) => {
  return (
    <>
      <Header />
      <main>{children}</main>
    </>
  )
}
CommonLayout.propTypes = {
  children: PropTypes.oneOfType([PropTypes.arrayOf(PropTypes.node), PropTypes.node])
}

export default CommonLayout
