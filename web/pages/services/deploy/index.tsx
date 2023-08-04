import {
  Alert,
  Button,
  Card,
  Label,
  Select,
  Spinner,
  Textarea,
  TextInput,
} from 'flowbite-react'
import Head from 'next/head'
import React, { useState } from 'react'
import { Formik, Form, Field } from 'formik'
import { createService } from '../../../utils/api'

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
            container_port: 80,
            env: '',
            autoscaling_metric: 'cpu',
            autoscaling_target: 80,
            min_scale: 0,
            max_scale: 3,
          }}
          onSubmit={async (values, { setSubmitting, resetForm }) => {
            console.log(values)
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

            const res = await createService({ ...values, env })
            const body = await res.json()
            if (res.ok) {
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
                <div className='sm:max-w-xs'>
                  <div className='mb-2 block'>
                    <Label htmlFor='container_port' value='Container port' />
                  </div>
                  <Field
                    as={TextInput}
                    name='container_port'
                    type='number'
                    placeholder='80'
                    required={true}
                  />
                </div>
                <div className='sm:max-w-md'>
                  <div className='grid md:grid-cols-2 md:gap-4'>
                    <div className='mb-4 md:mb-0'>
                      <div className='mb-2 block'>
                        <Label
                          htmlFor='autoscaling_metric'
                          value='Autoscaling metric'
                        />
                      </div>
                      <Field
                        as={Select}
                        name='autoscaling_metric'
                        required={true}
                      >
                        <option value='cpu'>CPU</option>
                        <option value='rps'>RPS</option>
                        <option value='memory'>Memory</option>
                      </Field>
                    </div>
                    <div>
                      <div className='mb-2 block'>
                        <Label
                          htmlFor='autoscaling_target'
                          value='Autoscaling target'
                        />
                      </div>
                      <Field
                        as={TextInput}
                        name='autoscaling_target'
                        type='number'
                        placeholder='80'
                        required={true}
                      />
                    </div>
                  </div>
                </div>
                <div className='max-w-md'>
                  <div className='grid md:grid-cols-2 md:gap-4'>
                    <div className='mb-4 md:mb-0'>
                      <div className='mb-2 block'>
                        <Label htmlFor='min_scale' value='Min scale' />
                      </div>
                      <Field
                        as={TextInput}
                        name='min_scale'
                        type='number'
                        placeholder='0'
                        required={true}
                      />
                    </div>
                    <div>
                      <div className='mb-2 block'>
                        <Label htmlFor='max_scale' value='Max scale' />
                      </div>
                      <Field
                        as={TextInput}
                        name='max_scale'
                        type='number'
                        placeholder='3'
                        required={true}
                      />
                    </div>
                  </div>
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
