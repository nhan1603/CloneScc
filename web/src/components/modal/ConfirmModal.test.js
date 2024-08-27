import ConfirmModal from './ConfirmModal'
import { render } from '@testing-library/react'

describe('ConfirmModal Component', () => {
  test('render ConfirmModal component successfully', () => {
    const { getByText } = render(
      <ConfirmModal title='Hello ConfirmModal' isOpen />
    )
    expect(getByText('Hello ConfirmModal')).toBeTruthy()
  })

  test('render ConfirmModal component successfully with delete button', () => {
    const { getByText } = render(
      <ConfirmModal title='Hello ConfirmModal with delete button' isOpen type='delete' />
    )
    expect(getByText('Hello ConfirmModal with delete button')).toBeTruthy()
  })
})
