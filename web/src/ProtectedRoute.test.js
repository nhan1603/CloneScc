import { render } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import ProtectedRoute from './ProtectedRoute'
import * as auth from 'data/auth'

describe('ProtectedRoute component', () => {
  test('render ProtectedRoute with children', () => {
    auth.getToken = () => true

    const props = {
      children: 'Hello World'
    }
    const { getByText } = render(
      <BrowserRouter>
        <ProtectedRoute children={props.children} />
      </BrowserRouter>
    )
    expect(getByText('Hello World')).toBeTruthy()
  })

  test('render ProtectedRoute with redirect', () => {
    auth.getToken = () => false

    const props = {
      children: 'Hello World'
    }
    const { queryByText } = render(
      <BrowserRouter>
        <ProtectedRoute children={props.children} />
      </BrowserRouter>
    )
    const text = queryByText('Hello World')
    expect(text).not.toBeInTheDocument()
  })
})
