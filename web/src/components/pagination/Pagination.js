/* eslint-disable react/prop-types */
import ReactPaginate from 'react-paginate'
import './Pagination.scss'
import classNames from 'classnames'

const PreviousLabel = <span>&lt;</span>
const NextLabel = <span>&gt;</span>
const BreakLabel = <span>&hellip;</span>

const Pagination = (props) => {
  return (
    <ReactPaginate
      containerClassName="pagination-container"
      pageClassName="pagination-page"
      pageLinkClassName="pagination-link"
      previousClassName={classNames('pagination-navigate pagination-previous', {
        disabled: props.pageCount > 0
      })}
      previousLinkClassName="pagination-link"
      nextClassName={classNames('pagination-navigate pagination-next', {
        disabled: props.pageCount <= 0
      })}
      nextLinkClassName="pagination-link"
      breakClassName="pagination-break"
      activeClassName="active"
      previousLabel={PreviousLabel}
      nextLabel={NextLabel}
      breakLabel={BreakLabel}
      marginPagesDisplayed={1}
      pageRangeDisplayed={3}
      disableInitialCallback
      {...props}
    />
  )
}

Pagination.displayName = 'Pagination'

export default Pagination
