import { render } from '@testing-library/react'
import FormFieldInput from './FormFieldInput'
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

describe('FormFieldInput Component', () => {
  test('render FormFieldInput component successfully with error message', () => {
    const { getByText } = render(
      <BrowserRouter>
        <FormFieldInput label='Label' name='Name' onChange={() => jest.fn()} />
      </BrowserRouter>
    )
    expect(getByText('Label')).toBeTruthy()
    expect(getByText('Error')).toBeTruthy()
  })
})
