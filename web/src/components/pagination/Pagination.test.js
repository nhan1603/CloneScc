import Pagination from './Pagination'
import { render } from '@testing-library/react'

describe('Pagination Component', () => {
  test('render Pagination component successfully', () => {
    const pageCount = 10
    const { getByText } = render(
      <Pagination pageCount={pageCount} page={0} />
    )
    expect(getByText(pageCount)).toBeTruthy()
  })
})
