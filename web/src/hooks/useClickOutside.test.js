import useClickOutside from './useClickOutside'
import { renderHook } from '@testing-library/react-hooks'

jest.mock('data/api', () => {
  return {
    ...jest.requireActual('data/api'),
    get: jest.fn()
  }
})

describe('useClickOutside', () => {
  it('useClickOutside with isOpen', async () => {
    // addEventListener will be called when isOpen = true
    window.addEventListener = jest.fn()
    const props = {
      ref: {
        current: {
          contains: () => false
        }
      },
      isOpen: true,
      handleClose: jest.fn()
    }
    renderHook(() => useClickOutside(props))
    expect(window.addEventListener).toHaveBeenCalled()
  })

  it('useClickOutside with isOpen = false', async () => {
    window.removeEventListener = jest.fn()
    const props = {
      ref: {
        current: {
          contains: () => false
        }
      },
      isOpen: false,
      handleClose: jest.fn()
    }
    renderHook(() => useClickOutside(props))
    expect(window.removeEventListener).toHaveBeenCalled()
  })
})
