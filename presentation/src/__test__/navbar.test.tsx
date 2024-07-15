
import { notificationsList, userDescriptor } from '@/components/mock'
import Navbar from '@/components/navbar/Navbar'
import '@testing-library/jest-dom'
import { render, screen } from '@testing-library/react'

describe('NavbarNull', () => {
    it('when user is null it render a div with two elements', () => {
        render(<Navbar user={null} notifications={[]} />)
        const list = screen.getByRole('list')
        const listItems = screen.getAllByRole('listitem')
        expect(list).toBeInTheDocument()
        listItems.forEach((li) => {
            expect(list).toContainElement(li)
        })
    })
})

describe('Navbar', () => {
    it('when user is not null it render the notifications and dropdown', () => {
        render(<Navbar user={userDescriptor} notifications={notificationsList} />)
        const list = screen.queryByRole('list')
        const listItems = screen.queryAllByRole('listitem')
        const buttons = screen.queryAllByRole('button')
        expect(list).toBeNull()
        expect(listItems).toHaveLength(0)
        expect(buttons).toHaveLength(2)
    })
})
