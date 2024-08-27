import { render } from '@testing-library/react'
import App from './App'

jest.mock('react-use-websocket', () => {
  return () => ({
    lastMessage: {
      data: '{"cctv_name":"CCTV name","floor_number":"G Floor","description":"Description","type":"Alert"}'
    }
  })
})

jest.mock('data/auth', () => ({
  getToken: () => true
}))

test('render App', () => {
  const { getByText } = render(<App />)
  expect(getByText('Oops! Something went wrong')).toBeTruthy()
})
