import classNames from 'classnames'
import CommonModal from './CommonModal'
import { CONFIRM_STYLE } from './modalStyles'
import './ConfirmModal.scss'
import PropTypes from 'prop-types'
import Button from 'components/button'

const { Header, Content, Footer } = CommonModal

const ConfirmModal = (props) => {
  const {
    title,
    isOpen,
    content,
    onClose,
    onConfirm,
    type,
    style,
    confirmLabel,
    cancelLabel,
    isConfirmDisabled,
    isCancelDisabled,
    ...modalProps
  } = props

  return (
    <CommonModal isOpen={isOpen} onClose={onClose} style={style} {...modalProps}>
      <Header onClose={onClose}>{title}</Header>
      <Content>{content}</Content>
      <Footer>
        <div className="footer-button">
          <Button buttonType="grey" onClick={onClose} disabled={isCancelDisabled}>
            {cancelLabel}
          </Button>
          <Button
            onClick={onConfirm}
            disabled={isConfirmDisabled}
            className={classNames(
              type === 'delete' && 'delete-button',
              'lm--button lm--button--solid'
            )}>
            {confirmLabel}
          </Button>
        </div>
      </Footer>
    </CommonModal>
  )
}

export default ConfirmModal
ConfirmModal.defaultProps = {
  confirmLabel: 'Confirm',
  cancelLabel: 'Cancel',
  type: 'confirm',
  style: CONFIRM_STYLE,
  isConfirmDisabled: false,
  isCancelDisabled: false
}

ConfirmModal.propTypes = {
  title: PropTypes.string,
  isOpen: PropTypes.bool,
  content: PropTypes.node,
  onClose: PropTypes.func,
  onConfirm: PropTypes.func,
  style: PropTypes.object,
  confirmLabel: PropTypes.string,
  cancelLabel: PropTypes.string,
  isConfirmDisabled: PropTypes.bool,
  isCancelDisabled: PropTypes.bool,
  type: PropTypes.string
}
