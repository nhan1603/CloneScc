import { render } from '@testing-library/react'
import { screen } from '@testing-library/dom'
import userEvent from '@testing-library/user-event'
import { BrowserRouter } from 'react-router-dom'
import PaginationSummary from './PaginationSummary'

jest.mock('hooks/usePremiseParam', () => {
  return () => ({
    premiseId: '123',
    withPremiseParam: () => {
      return () => {
        return '/hello'
      }
    }
  })
})

describe('PaginationSummary Component', () => {
  test('render PaginationSummary with default value', () => {
    const { getAllByText } = render(
      <BrowserRouter>
        <PaginationSummary
          page={1}
          totalCount={30}
          pageCount={10}
          onPageChange={() => {}}
        />
      </BrowserRouter>
    )
    expect(getAllByText('1')[0]).toBeTruthy()
  })
})
