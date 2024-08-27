import Modal from 'react-modal'
import { MODAL_STYLE } from './modalStyles'
import './CommonModal.scss'
import PropTypes from 'prop-types'

const Header = ({ children, onClose, headerDescription }) => (
  <div className="modal-header">
    <div className="modal-header--content">
      <div className="cancel-button" onClick={onClose}>
        X
      </div>
      <div className="title">
        {children}
        {!!headerDescription && <div className="description">{headerDescription}</div>}
      </div>
    </div>
  </div>
)

const Content = ({ children }) => <div className="modal-content">{children}</div>

const Footer = ({ children }) => <div className="modal-footer">{children}</div>

const CommonModal = ({ isOpen, children, style = MODAL_STYLE, ...modalProps }) => {
  return (
    <Modal isOpen={isOpen} style={style} ariaHideApp={false} {...modalProps}>
      <div className="modal-container">{children}</div>
    </Modal>
  )
}

CommonModal.Header = Header
CommonModal.Content = Content
CommonModal.Footer = Footer

Header.propTypes = {
  children: PropTypes.node,
  onClose: PropTypes.func,
  headerDescription: PropTypes.string
}

Content.propTypes = {
  children: PropTypes.node
}

Footer.propTypes = {
  children: PropTypes.node
}

CommonModal.propTypes = {
  children: PropTypes.node,
  isOpen: PropTypes.bool,
  style: PropTypes.object
}
export default CommonModal
