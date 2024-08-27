import PropsTypes from 'prop-types'
import './PaginationSummary.scss'
import Pagination from 'components/pagination/Pagination'
import { LIMIT } from 'utils/contants'

const PaginationSummary = ({ page, totalCount, pageCount, onPageChange, data, ...otherProps }) => {
  const numOfItems = data.length
  return (
    <div className="pagination-summary">
      <div className="pagination-summary__display-string">
        {page === 1 ? 1 : LIMIT * page + 1 - LIMIT} -{' '}
        {numOfItems === LIMIT ? numOfItems * page : totalCount} of {totalCount}
      </div>

      <Pagination
        pageCount={pageCount}
        forcePage={page - 1}
        onPageChange={onPageChange}
        {...otherProps}
      />
    </div>
  )
}

PaginationSummary.defaultProps = {
  data: []
}

PaginationSummary.propTypes = {
  page: PropsTypes.number,
  totalCount: PropsTypes.number,
  pageCount: PropsTypes.number,
  onPageChange: PropsTypes.func,
  data: PropsTypes.array
}
export default PaginationSummary
