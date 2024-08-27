import { render } from '@testing-library/react'
import Button from './Button'

describe('Button tests', () => {
  it('should display valid button', () => {
    const { getByText, container } = render(
      <Button buttonSize="small" buttonType="solid">
        Button test
      </Button>
    )
    expect(getByText('Button test')).toBeInTheDocument()
    expect(container.getElementsByClassName('lm--button')).toBeTruthy()
    expect(container.getElementsByClassName('lm--button--small')).toBeTruthy()
    expect(container.getElementsByClassName('lm--button--solid')).toBeTruthy()
  })
})
