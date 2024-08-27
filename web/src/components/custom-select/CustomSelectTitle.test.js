import { render } from '@testing-library/react'
import CustomSelectTitle from './CustomSelectTitle'
import { BrowserRouter } from 'react-router-dom'

describe('CustomSelectTitle Component', () => {
  test('render CustomSelectTitle with default value', () => {
    const { getByText } = render(
      <BrowserRouter>
        <CustomSelectTitle
          valueContainerLabel="valueContainerLabel"
          subOption="Sub Option"
          value="value"
          options={['value']}
        />
      </BrowserRouter>
    )
    expect(getByText('All premises')).toBeTruthy()
  })
})
