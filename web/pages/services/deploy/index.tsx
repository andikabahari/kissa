import {
  Alert,
  Button,
  Card,
  Label,
  Spinner,
  Textarea,
  TextInput,
} from 'flowbite-react'
import Head from 'next/head'
import React, { useState } from 'react'
import { Formik, Form, Field } from 'formik'

export default function Deploy() {
  const [alertSuccessOpen, setAlertSuccessOpen] = useState<boolean>(false)
  const [alertFailureOpen, setAlertFailureOpen] = useState<boolean>(false)
  const [alertFailureMsg, setAlertFailureMsg] = useState<string>('')

  return (
    <>
      <Head>
        <title>Deploy a New Service</title>
      </Head>
      <Alert
        color='success'
        className={alertSuccessOpen ? 'mb-5' : 'hidden'}
        onDismiss={() => setAlertSuccessOpen(false)}
      >
        <span>Service deployed successfully.</span>
      </Alert>
      <Alert
        color='failure'
        className={alertFailureOpen ? 'mb-5' : 'hidden'}
        onDismiss={() => setAlertFailureOpen(false)}
      >
        <span>{alertFailureMsg}</span>
      </Alert>
      <Card>
        <h5 className='text-2xl font-bold tracking-tight text-gray-900 dark:text-white'>
          New service
        </h5>
        <Formik
          initialValues={{
            name: '',
            image: '',
            containerPort: 0,
            env: '',
          }}
          onSubmit={async (values, { setSubmitting, resetForm }) => {
            const env = values.env
              .split('\n')
              .map((val) => {
                const pairs = val.split('=')
                return {
                  name: pairs[0] ?? '',
                  value: pairs[1] ?? '',
                }
              })
              .filter(({ name }) => Boolean(name))

            const res = await fetch('/api/services', {
              method: 'POST',
              headers: { 'Content-Type': 'application/json' },
              body: JSON.stringify({ ...values, env }),
            })
            const body = await res.json()

            if (body.code === 200) {
              setAlertSuccessOpen(true)
              resetForm()
            } else {
              setAlertFailureOpen(true)
              setAlertFailureMsg(body.message)
            }

            setSubmitting(false)
          }}
        >
          {({ isSubmitting }) => (
            <Form>
              <div className='flex flex-col gap-4 mb-4'>
                <div>
                  <div className='mb-2 block'>
                    <Label htmlFor='name' value='Name' />
                  </div>
                  <Field
                    as={TextInput}
                    name='name'
                    type='text'
                    placeholder='hello-world'
                    required={true}
                  />
                </div>
                <div>
                  <div className='mb-2 block'>
                    <Label htmlFor='image' value='Image' />
                  </div>
                  <Field
                    as={TextInput}
                    name='image'
                    type='text'
                    placeholder='gcr.io/knative-samples/helloworld-go'
                    required={true}
                  />
                </div>
                <div className='max-w-xs'>
                  <div className='mb-2 block'>
                    <Label htmlFor='containerPort' value='Container port' />
                  </div>
                  <Field
                    as={TextInput}
                    name='containerPort'
                    type='number'
                    placeholder='8080'
                    required={true}
                  />
                </div>
                <div>
                  <div className='mb-2 block'>
                    <Label htmlFor='env' value='Environment variables' />
                  </div>
                  <Field
                    as={Textarea}
                    name='env'
                    placeholder='TARGET=World'
                    className='h-24'
                  />
                </div>
              </div>
              <Button
                type='submit'
                gradientDuoTone='cyanToBlue'
                disabled={isSubmitting}
              >
                {isSubmitting ? (
                  <>
                    <span className='mr-2'>
                      <Spinner size='sm' light={true} />
                    </span>
                    Loading...
                  </>
                ) : (
                  'Submit'
                )}
              </Button>
            </Form>
          )}
        </Formik>
      </Card>
    </>
  )
}
