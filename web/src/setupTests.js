// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom'
jest.mock('react-use-websocket', () => {
  return () => ({
    lastMessage: {
      data: '{"cctv_name":"CCTV name","floor_number":"G Floor","description":"Description","type":"Alert"}'
    }
  })
})
