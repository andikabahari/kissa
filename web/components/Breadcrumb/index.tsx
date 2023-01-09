import { Breadcrumb as FlowbiteBreadcrumb } from 'flowbite-react'
import { useRouter } from 'next/router'
import React from 'react'

export default function Breadcrumb() {
  const router = useRouter()
  return (
    <FlowbiteBreadcrumb className='mb-8'>
      {router.pathname.split('/').map((val, idx) => {
        if (idx === 0) {
          return <FlowbiteBreadcrumb.Item>Home</FlowbiteBreadcrumb.Item>
        }

        if (val) {
          return <FlowbiteBreadcrumb.Item>{val}</FlowbiteBreadcrumb.Item>
        }
      })}
    </FlowbiteBreadcrumb>
  )
}
