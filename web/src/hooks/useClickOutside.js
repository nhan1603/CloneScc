import { useEffect } from 'react'

const useClickOutside = ({ ref, isOpen, handleClose }) => {
  useEffect(() => {
    const listener = (event) => {
      if (isOpen && ref.current && !ref.current.contains(event.target)) {
        handleClose()
      }
    }
    isOpen && window.addEventListener('mouseup', listener)
    !isOpen && window.removeEventListener('mouseup', listener)
    return () => {
      window.removeEventListener('mouseup', listener)
    }
  }, [isOpen, ref, handleClose])
}

export default useClickOutside
