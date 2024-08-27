import CreateRequestModal from './CreateRequestModal'
import { fireEvent, render } from '@testing-library/react'
import { act } from 'react-dom/test-utils'

describe('CreateRequestModal Component', () => {
  test('render CreateRequestModal component successfully', async () => {
    const callbackFn = jest.fn()
    const { getByText } = render(
      <CreateRequestModal title='Hello CreateRequestModal' isOpen onSubmitCallback={callbackFn} />
    )
    act(() => {
      /* fire events that update state */
      fireEvent.click(getByText('Send Request'))
    })
    expect(getByText('CREATE REQUEST')).toBeTruthy()
  })

  test('render CreateRequestModal component is close', async () => {
    const callbackFn = jest.fn()
    const { queryByText } = render(
      <CreateRequestModal title='Hello CreateRequestModal' isOpen={false} onSubmitCallback={callbackFn} />
    )
    const titleModal = queryByText('CREATE REQUEST')
    expect(titleModal).not.toBeInTheDocument()
  })

  test('render CreateRequestModal component successfully and close this', async () => {
    const callbackFn = jest.fn()
    const { getByText } = render(
      <CreateRequestModal title='Hello CreateRequestModal' isOpen onSubmitCallback={callbackFn} />
    )
    act(() => {
      /* fire events that update state */
      fireEvent.click(getByText('Cancel'))
    })
    expect(getByText('CREATE REQUEST')).toBeTruthy()
  })

  test('render CreateRequestModal component send request and show confirm modal', async () => {
    // TODO: will update this
    const { getByText } = render(
      <CreateRequestModal title='Hello CreateRequestModal' isOpen />
    )
    act(() => {
      /* fire events that update state */
      fireEvent.click(getByText('Cancel'))
    })
    expect(getByText('CREATE REQUEST')).toBeTruthy()
  })
})
