
import { notificationsList } from '@/components/mock'
import Notifications from '@/components/notifications/Notifications'
import '@testing-library/jest-dom'
import { fireEvent, render, screen } from '@testing-library/react'

describe('NotificationsEmpty', () => {
    it('when button is clicked, it renders a list with one item that has a paragraph, when clicked again, it hides the list', () => {
        render(<Notifications notifications={[]} />)
        let list = screen.queryByRole('list')
        let listItem = screen.queryByRole('listitem')
        let paragraph = screen.queryByRole('paragraph')
        const button = screen.getByRole('button')
        expect(list).toBeNull()
        expect(listItem).toBeNull()
        expect(paragraph).toBeNull()
        expect(button).toBeInTheDocument()
        fireEvent(button,
            new MouseEvent('click', {
                bubbles: true,
                cancelable: true,
            })
        )
        list = screen.getByRole('list')
        listItem = screen.getByRole('listitem')
        paragraph = screen.getByRole('paragraph')
        expect(list).toBeInTheDocument()
        expect(paragraph).toBeInTheDocument()
        expect(listItem).toBeInTheDocument()
        expect(listItem).toContainElement(paragraph)
        expect(list).toContainElement(listItem)
        fireEvent(button,
            new MouseEvent('click', {
                bubbles: true,
                cancelable: true,
            })
        )
        expect(list).not.toBeInTheDocument()
        expect(paragraph).not.toBeInTheDocument()
        expect(listItem).not.toBeInTheDocument()
    })
})

describe('Notifications', () => {
    it('when button is clicked and a mock with correct type exists and its length is greater than zero,it renders the mock as a list', () => {
        render(<Notifications notifications={notificationsList} />)
        let list = screen.queryByRole('list')
        let listItem = screen.queryAllByRole('listitem')
        const button = screen.getByRole('button')
        expect(list).toBeNull()
        expect(listItem).toHaveLength(0)
        expect(button).toBeInTheDocument()
        fireEvent(button,
            new MouseEvent('click', {
                bubbles: true,
                cancelable: true,
            })
        )
        list = screen.getByRole('list')
        listItem = screen.getAllByRole('listitem')
        expect(listItem).toHaveLength(notificationsList.length)
        expect(list).toBeInTheDocument()
        listItem.forEach((li)=>{
            expect(li).toBeInTheDocument()
            expect(list).toContainElement(li)
        })
        fireEvent(button,
            new MouseEvent('click', {
                bubbles: true,
                cancelable: true,
            })
        )
        expect(list).not.toBeInTheDocument()
        listItem.forEach((li) => expect(li).not.toBeInTheDocument())
    })
})