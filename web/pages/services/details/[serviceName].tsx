import Link from 'next/link'
import { Badge, Button, Modal, Table } from 'flowbite-react'
import Head from 'next/head'
import { useRouter } from 'next/router'
import { getRevisions } from '../../../utils/api'
import { useEffect, useState } from 'react'

export default function Home() {
  const router = useRouter()

  const {
    data: revisions,
    error,
    isLoading,
  } = getRevisions(router.query.serviceName as string)
  const [revisionToDelete, setRevisionToDelete] = useState<string>('')
  const [deleteModal, setDeleteModal] = useState<string>('')

  return (
    <>
      <Head>
        <title>{router.query.serviceName}</title>
      </Head>
      <div className='overflow-y-auto max-h-[576px] mt-8'>
        <Table hoverable={true}>
          <Table.Head>
            <Table.HeadCell>Name</Table.HeadCell>
            <Table.HeadCell>Ready</Table.HeadCell>
            <Table.HeadCell>Actual replicas</Table.HeadCell>
            <Table.HeadCell>Derised replicas</Table.HeadCell>
            <Table.HeadCell>Created at</Table.HeadCell>
            <Table.HeadCell>
              <span className='sr-only'>Details</span>
            </Table.HeadCell>
          </Table.Head>
          <Table.Body className='divide-y'>
            {isLoading ? (
              <Table.Row>
                <Table.Cell colSpan={6} className='bg-white'>
                  feching data...
                </Table.Cell>
              </Table.Row>
            ) : error ? (
              <Table.Row>
                <Table.Cell colSpan={6} className='bg-white'>
                  {error.message}
                </Table.Cell>
              </Table.Row>
            ) : (
              revisions.map((val: any, idx: number) => (
                <Table.Row
                  key={idx}
                  className='bg-white dark:border-gray-700 dark:bg-gray-800'
                >
                  <Table.Cell className='whitespace-nowrap font-medium text-gray-900 dark:text-white'>
                    {val.name}
                  </Table.Cell>
                  <Table.Cell>
                    <div className='inline-block'>
                      {val.ready ? (
                        <Badge color='success'>True</Badge>
                      ) : (
                        <Badge color='failure'>False</Badge>
                      )}
                    </div>
                  </Table.Cell>
                  <Table.Cell>{val.actual_replicas}</Table.Cell>
                  <Table.Cell>{val.desired_replicas}</Table.Cell>
                  <Table.Cell>{val.created_at}</Table.Cell>
                  <Table.Cell>
                    <Button
                      size='xs'
                      color='failure'
                      onClick={() => {
                        setRevisionToDelete(val.name)
                        setDeleteModal('default')
                      }}
                    >
                      Delete
                    </Button>
                  </Table.Cell>
                </Table.Row>
              ))
            )}
          </Table.Body>
        </Table>
      </div>

      <Modal
        show={deleteModal === 'default'}
        onClose={() => setDeleteModal('')}
      >
        <Modal.Body>
          <div className='space-y-6'>
            <p className='text-lg'>Are you sure?</p>
            <p className='text-base leading-relaxed text-gray-500'>
              You are about to delete revision{' '}
              <span className='font-medium'>"{revisionToDelete}"</span>. Once
              deleted, you cannot rollback service to this revision.
            </p>
            <div className='flex flex-wrap gap-2'>
              <Button
                size='sm'
                color='gray'
                onClick={() => {
                  setRevisionToDelete('')
                  setDeleteModal('')
                }}
              >
                No, cancel
              </Button>
              <Button
                size='sm'
                color='failure'
                onClick={() => {
                  setRevisionToDelete('')
                  setDeleteModal('')
                  // TODO: perform deletion -> deleteRevision(revisionToDelete)
                }}
              >
                Yes, delete
              </Button>
            </div>
          </div>
        </Modal.Body>
      </Modal>
    </>
  )
}
