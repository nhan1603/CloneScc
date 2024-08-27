import { render } from '@testing-library/react'
import Toast, { addToast, addSuccessToast, addErrorToast } from './Toast'
import { BrowserRouter } from 'react-router-dom'

jest.mock('react-hook-form', () => {
  return {
    ...jest.requireActual('react-hook-form'),
    useController: () => {
      return {
        field: {
          value: 'hello'
        },
        fieldState: {
          error: {
            message: 'Error'
          }
        }
      }
    }
  }
})

jest.mock('hooks/usePremiseParam', () => {
  return () => ({
    withPremiseParam: () => {
      return () => {
        return '/hello'
      }
    }
  })
})

describe('Toast Component', () => {
  test('render Toast component successfully', () => {
    const { getByText } = render(
      <BrowserRouter>
        <Toast title='Title Toast' message='Toast success' status='success' />
      </BrowserRouter>
    )
    addSuccessToast({ message: 'Toast success' })
    expect(getByText('Toast success')).toBeTruthy()
  })

  test('render Toast component successfully status=error', () => {
    const { getByText } = render(
      <BrowserRouter>
        <Toast title='Title Toast' message='Toast error' status='error' date='Today' />
      </BrowserRouter>
    )
    addErrorToast({ message: 'Toast error' })
    expect(getByText('Toast error')).toBeTruthy()
  })

  test('render Toast component successfully status=info', () => {
    const { getByText } = render(
      <BrowserRouter>
        <Toast title='Title Toast' message='Toast info' status='info' />
      </BrowserRouter>
    )
    addToast({ message: 'Toast info' })
    expect(getByText('Toast info')).toBeTruthy()
  })
})
