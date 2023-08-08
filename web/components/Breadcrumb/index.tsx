import { Breadcrumb as FlowbiteBreadcrumb } from 'flowbite-react'
import { useRouter } from 'next/router'
import React from 'react'

export default function Breadcrumb() {
  const router = useRouter()
  const pathnames = router.pathname.split('/').map((segment) => {
    for (const key of Object.keys(router.query)) {
      if (segment.includes(key)) return router.query[key]
    }
    return segment
  })
  return (
    <FlowbiteBreadcrumb className='mb-8'>
      {pathnames.map((val, idx) => {
        if (idx === 0) {
          return (
            <FlowbiteBreadcrumb.Item key={idx}>Home</FlowbiteBreadcrumb.Item>
          )
        }
        if (val) {
          return (
            <FlowbiteBreadcrumb.Item key={idx}>{val}</FlowbiteBreadcrumb.Item>
          )
        }
      })}
    </FlowbiteBreadcrumb>
  )
}
