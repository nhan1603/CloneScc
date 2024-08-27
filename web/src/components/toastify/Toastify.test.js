import { render } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import Toastify from './Toastify'

describe('Toastify', () => {
  it('should render Toastify successfully', async () => {
      process.env.REACT_APP_SOCKET_API_ENDPOINT = '/hello'
      const { baseElement } = render(
        <BrowserRouter>
          <Toastify />
        </BrowserRouter>
      )
      expect(baseElement).toBeTruthy()
  });
});
