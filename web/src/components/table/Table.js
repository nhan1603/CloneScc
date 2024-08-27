/* eslint-disable react/prop-types */
import React, { useMemo } from 'react'
import PropTypes from 'prop-types'
import ReactTable from 'react-table'
import classNames from 'classnames'
import './Table.scss'

const Table = (props) => {
  const { className, columns, ...otherProps } = props
  const { loading, sortable } = otherProps
  const finalColumns = useMemo(() => {
    let retColumn
    return columns.map((column) => {
      retColumn = column
      if (column.alignRight) {
        retColumn = {
          ...column,
          className: classNames({
            [column.className]: column.className,
            '-right': true
          }),
          headerClassName: classNames({
            [column.headerClassName]: column.headerClassName,
            '-right': true
          }),
          Header: (props) => (
            <span>
              {typeof column.Header === 'function' ? column.Header(props) : column.Header}
            </span>
          ),
          Cell: (props) => (
            <span>
              {typeof column.Cell === 'function' ? column.Cell(props) : column.Cell || props.value}
            </span>
          )
        }
      }
      if (sortable === false || column.sortable === false) {
        retColumn = {
          ...retColumn,
          headerClassName: classNames({
            [retColumn.headerClassName]: retColumn.headerClassName,
            '-non-sortable': true
          })
        }
      }
      return retColumn
    })
  }, [columns, sortable])

  return (
    <ReactTable
      className={`${className} ${loading ? '-loading' : ''}`}
      columns={finalColumns}
      resizable={false}
      showPagination={false}
      minRows={1}
      noDataText="No records found"
      {...otherProps}
    />
  )
}

Table.displayName = 'Table'

Table.propTypes = {
  /**
   * Class name
   */
  className: PropTypes.string,
  /**
   * Loading flag
   */
  loading: PropTypes.bool,
  /**
   * Sortable flag
   */
  sortable: PropTypes.bool,
  /**
   * Columns
   */
  columns: PropTypes.arrayOf(
    PropTypes.shape({
      /**
       * Flag to indicate if header and cell text should be aligned right
       */
      alignRight: PropTypes.bool
    })
  )
}

Table.defaultProps = {
  className: ''
}

export default Table
