import CommonModal from './CommonModal'
import { render } from '@testing-library/react'

describe('CommonModal Component', () => {
  test('render CommonModal component successfully', () => {
    const { getByText } = render(
      <>
        <CommonModal.Content />
        <CommonModal.Footer />
        <CommonModal children='Hello CommonModal' isOpen={true} onClose={jest.fn()} headerDescription='header Description' />
      </>
    )
    expect(getByText('Hello CommonModal')).toBeTruthy()
  })

  test('render CommonModal component successfully', () => {
    const { getByText } = render(
      <CommonModal.Header children='Hello Header' onClose={jest.fn()} headerDescription='header Description' />
    )
    expect(getByText('Hello Header')).toBeTruthy()
    expect(getByText('header Description')).toBeTruthy()
  })

  test('render CommonModal component successfully with no headerDescription', () => {
    const { getByText } = render(
      <CommonModal.Header children='header with no Description' onClose={jest.fn()} headerDescription={null} />
    )
    expect(getByText('header with no Description')).toBeTruthy()
  })
})
