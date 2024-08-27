const DEFAULT_STYLE = {
  content: {
    top: '50%',
    left: '50%',
    right: 'auto',
    bottom: 'auto',
    transform: 'translate(-50%, -50%)',
    width: '100%'
  },
  overlay: {
    backgroundColor: 'rgba(0, 0, 0, 0.7)'
  }
}

export const MODAL_STYLE = {
  content: {
    ...DEFAULT_STYLE.content,
    maxWidth: '800px',
    maxHeight: '100vh',
    padding: '0'
  },
  overlay: {
    ...DEFAULT_STYLE.overlay
  }
}

export const CONFIRM_STYLE = {
  ...MODAL_STYLE,
  content: {
    ...MODAL_STYLE.content,
    maxWidth: '500px'
  }
}
