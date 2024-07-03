import { createRootRouteWithContext } from '@tanstack/react-router'
import Root from '../components/Root'
import { fetchUserDescriptor } from '../lib/fetchUserDescriptor'
import { fetchNotifications } from '../lib/fetchNotifications'
import { Notifications } from '../lib/types'
import { hasPermission, isLoggedIn } from '../lib/is_logged_in'
export const Route = createRootRouteWithContext<{
  hasPermission: typeof hasPermission
  fetchUserDescriptor: typeof fetchUserDescriptor
  fetchNotifications: typeof fetchNotifications
  isLoggedIn: typeof isLoggedIn
}>()({
  loader: async () => {
    let descriptor = null
    let id = null
    let notifications: Notifications[] = []
    try {
      descriptor = await fetchUserDescriptor()
      if (descriptor) {
        id = `notifications:${descriptor.email}`
        for (let i = 0; i < descriptor.grantedAuthorities.length; i++) {
          if (i === descriptor.grantedAuthorities.length - 1) {
            id += `?notifications:${descriptor.grantedAuthorities[i].enterpriseID}`
          } else {
            id += `?notifications:${descriptor.grantedAuthorities[i].enterpriseID}?`
          }
        }
        notifications = await fetchNotifications(id)
      }
      return { descriptor, notifications, id }
    } catch (e) {

      return { descriptor, notifications, id }
    }
  },
  component: Root,
})
