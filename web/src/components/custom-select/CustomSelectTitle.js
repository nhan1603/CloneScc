import React, { useEffect, useState } from 'react'
import CustomSelect from './CustomSelect'
import PropTypes from 'prop-types'
import useGet from 'hooks/useGet'
import { GET_PREMISES_URL } from 'utils/contants'
import usePremiseParam from 'hooks/usePremiseParam'

const DEFAULT_OPTION = {
  value: 'all',
  label: 'All',
  subOption: 'All premises'
}
const CustomSelectTitle = ({ handlePremisesChange }) => {
  const [premises, setPremises] = useState([])
  const { premiseId, setPremiseParam } = usePremiseParam()

  const { response } = useGet(GET_PREMISES_URL)
  const [premisesShowing, setPremisesShowing] = useState({})

  const items = response?.data?.items ?? []
  useEffect(() => {
    const matchedPremise = items?.find((item) => item.id === Number(premiseId))
    if (matchedPremise) {
      setPremisesShowing({
        value: matchedPremise.id,
        label: matchedPremise.name,
        subOption: matchedPremise.location
      })
    } else {
      setPremisesShowing(DEFAULT_OPTION)
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [JSON.stringify(items)])

  const getRemainingPremises = () => {
    return premises?.filter((premise) => premisesShowing.value !== premise?.value)
  }

  const setPremisesOptions = () => {
    const allPremisesOption = items?.map((item) => ({
      value: item.id,
      label: item.name,
      subOption: item.location
    }))
    setPremises([DEFAULT_OPTION, ...allPremisesOption])
  }

  const handleOnChange = (value) => {
    setPremisesShowing(value)
    handlePremisesChange?.(value)
    setPremiseParam(value.value)
  }

  useEffect(() => {
    items?.length > 0 && setPremisesOptions()
  }, [JSON.stringify(items)])

  return (
    <CustomSelect
      subOption={premisesShowing.subOption}
      options={getRemainingPremises()}
      value={premisesShowing}
      onChange={handleOnChange}
      isSearchable={false}
    />
  )
}

CustomSelectTitle.propTypes = {
  handlePremisesChange: PropTypes.func
}

export default React.memo(CustomSelectTitle)
