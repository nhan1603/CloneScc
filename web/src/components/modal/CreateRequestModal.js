import Button from 'components/button'
import CommonModal from './CommonModal'
import { FormProvider, useForm } from 'react-hook-form'
import FormTextArea from 'components/form/form-text-area'
import './CreateRequestModal.scss'
import FormBaseSelect from 'components/form/form-base-select'
import { zodResolver } from '@hookform/resolvers/zod'
import { useEffect, useState } from 'react'
import ConfirmModal from './ConfirmModal'
import PropTypes from 'prop-types'
import requestSchema from 'schemas/requestSchema'
import usePost from '../../hooks/usePost'
import { REQUESTS_URL } from 'utils/contants'
import useGet from 'hooks/useGet'
import { USERS_URL } from '../../utils/contants'
import { getCurrentUser } from 'data/auth'

const CREATE_REQUEST_TYPE = 'NEW'

const CreateRequestModal = ({ isOpen, onClose, onSuccess, onError, alertDetails }) => {
  const { response } = useGet(USERS_URL)
  const allUsers = response?.data?.items ?? []
  const allAdmin = allUsers.filter((user) => user.role === 'OPERATION_USER')
  const otherUsers = allUsers.filter((user) => user.role !== 'OPERATION_USER')

  const options = otherUsers.map((user) => ({ label: user.name, value: user.id }))

  const defaultValues = {
    assignee: { label: '', value: '' },
    message: ''
  }
  const form = useForm({
    defaultValues,
    resolver: zodResolver(requestSchema),
    mode: 'all'
  })

  const { post, isLoading: isLoadingPost } = usePost()

  const [isOpenConfirm, setIsOpenConfirm] = useState(false)

  const handleSendRequest = async () => {
    const values = form.getValues()
    const currAdmin = getCurrentUser()
    const adminId = allAdmin.find((admin) => admin.email === currAdmin.email)?.id
    const payload = {
      alertID: alertDetails.id,
      requestBy: adminId,
      assignedUserID: values.assignee.value,
      content: values.message,
      type: CREATE_REQUEST_TYPE
    }
    await post(REQUESTS_URL, payload, onSuccess, onError)
  }

  useEffect(() => {
    if (!isOpen) form.reset()
  }, [isOpen])

  const onCloseModal = () => {
    // Todo: investigate why form.formState.isDirty doesn't work as expected
    const currentValues = form.getValues()
    if (JSON.stringify(currentValues) !== JSON.stringify(defaultValues)) {
      setIsOpenConfirm(true)
    } else {
      onClose?.()
    }
  }
  return (
    <>
      <CommonModal isOpen={isOpen}>
        <CommonModal.Header onClose={onCloseModal}>CREATE REQUEST</CommonModal.Header>
        <CommonModal.Content>
          <FormProvider {...form}>
            <div className="create-request-modal">
              <div className="create-request-modal__details">
                <div className="details-label">
                  <label>Detailed Information</label>
                </div>
                <div className="details-description">
                  <div className="title">Type</div>
                  <div>{alertDetails?.type}</div>
                  <div className="title">Date</div>
                  <div>{alertDetails?.date}</div>
                  <div className="title">Premises</div>
                  <div>
                    <div>{alertDetails?.premise?.name}</div>
                    <div>{alertDetails?.premise?.location}</div>
                  </div>
                  <div className="title">Camera ID</div>
                  <div>{alertDetails?.cameraId}</div>
                </div>
              </div>
              <FormTextArea label="Message" name="message" grid />
              <FormBaseSelect label="Assignee" name="assignee" options={options} required grid />
            </div>
          </FormProvider>
        </CommonModal.Content>
        <CommonModal.Footer>
          <div className="request-button-group">
            <Button buttonType="grey" disabled={isLoadingPost} onClick={onCloseModal}>
              Cancel
            </Button>
            <Button
              buttonType="solid"
              disabled={isLoadingPost}
              onClick={() => {
                form.handleSubmit(handleSendRequest)()
              }}>
              {isLoadingPost ? 'Sending...' : 'Send Request'}
            </Button>
          </div>
        </CommonModal.Footer>
      </CommonModal>
      <ConfirmModal
        isOpen={isOpenConfirm}
        onClose={() => {
          setIsOpenConfirm(false)
        }}
        onConfirm={() => {
          setIsOpenConfirm(false)
          onClose?.()
        }}
        title="Cancel Request"
        content="Are you sure you want to cancel the request?"
        cancelLabel="No"
        confirmLabel="Yes"
      />
    </>
  )
}

export default CreateRequestModal

CreateRequestModal.propTypes = {
  isOpen: PropTypes.bool,
  onSuccess: PropTypes.func,
  onClose: PropTypes.func,
  onError: PropTypes.func,
  alertDetails: PropTypes.object
}
