import { CustomSelect, CustomSelectTitle } from 'components/custom-select'
import CommonLayout from 'components/layouts'
import Loading from 'components/loading'
import BodyLayout from 'components/layouts/BodyLayout'
import useGet from 'hooks/useGet'
import { useEffect, useState } from 'react'
import { GET_DEVICES_URL } from 'utils/contants'
import usePremiseParam from 'hooks/usePremiseParam'
import './Dashboard.scss'

const NUMBER_OF_CAMERAS = 4

const Dashboard = () => {
  const [camerasShowing, setCamerasShowing] = useState([])
  const [allCamerasOption, setAllCamerasOption] = useState([])
  const [param, setParam] = useState({})
  const { premiseId } = usePremiseParam()
  const { response, isLoading } = useGet(GET_DEVICES_URL, param)

  const getRemainingCameras = (allCameras, selectedCameras) => {
    const selectedCameraIds = selectedCameras?.map((item) => item?.value)
    return allCameras?.filter((camera) => !selectedCameraIds?.includes(camera?.value))
  }

  const handleChangeCameras = (index) => (cameraShowing) => {
    const newCamerasShowing = [...camerasShowing]
    newCamerasShowing[index] = cameraShowing
    setCamerasShowing(newCamerasShowing)
  }

  const items = response?.data?.items
  const setCamerasOptions = () => {
    const allCameras = items?.map((item) => ({
      value: item.id,
      label: item.premiseName + ' - ' + item.deviceName,
      deviceURL: `${process.env.REACT_APP_API_ENDPOINT}/api/public/v1/camera/${item.deviceCode}`
    }))
    setAllCamerasOption([...allCameras])
    setCamerasShowing([...allCameras].slice(0, NUMBER_OF_CAMERAS))
  }

  const handlePremisesChange = (dataPremise) => {
    setParam(dataPremise?.value !== 'all' ? { premiseID: dataPremise.value } : {})
  }

  useEffect(() => {
    items?.length > 0 && setCamerasOptions()
  }, [JSON.stringify(items)])

  useEffect(() => {
    premiseId && handlePremisesChange({ value: premiseId })
  }, [premiseId])

  return (
    <CommonLayout>
      <BodyLayout leftContent={<CustomSelectTitle handlePremisesChange={handlePremisesChange} />}>
        {isLoading ? (
          <Loading />
        ) : (
          <div className="camera">
            {camerasShowing.map((camera, index) => {
              return (
                <div className="camera__item" key={index}>
                  <div className="camera__content">
                    <iframe
                      width="100%"
                      height="400"
                      src={camera.deviceURL}
                      frameBorder="0"
                      allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                      allowFullScreen
                      title={`SCC Camera ${camera.id}`}
                    />
                  </div>
                  <CustomSelect
                    options={getRemainingCameras(allCamerasOption, camerasShowing)}
                    value={camerasShowing[index]}
                    onChange={handleChangeCameras(index)}
                    isSearchable={false}
                  />
                </div>
              )
            })}
          </div>
        )}
      </BodyLayout>
    </CommonLayout>
  )
}

export default Dashboard
