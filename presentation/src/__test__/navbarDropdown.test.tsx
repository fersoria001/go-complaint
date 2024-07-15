
import { userDescriptor } from '@/components/mock'
import NavbarDropdown from '@/components/navbar/NavbarDropdown'
import '@testing-library/jest-dom'
import { fireEvent, render, screen } from '@testing-library/react'

describe('NavbarDropdown', () => {
    it('when button is clicked it renders a div with a list, when an item from the list is clicked, it hides the list.', () => {
        render(<NavbarDropdown user={userDescriptor} />)
        let list = screen.queryByRole('list')
        let listItems = screen.queryAllByRole('listitem')
        let links = screen.queryAllByRole('link')
        const button = screen.getByRole('button')
        expect(list).toBeNull()
        expect(listItems).toHaveLength(0)
        expect(links).toHaveLength(0)
        expect(button).toBeInTheDocument()
        fireEvent(button,
            new MouseEvent('click', {
                bubbles: true,
                cancelable: true,
            })
        )
        list = screen.getByRole('list')
        listItems = screen.getAllByRole('listitem')
        links = screen.getAllByRole('link')
        expect(list).toBeInTheDocument()
        listItems.forEach((li) => {
            expect(li).toBeInTheDocument()
            expect(list).toContainElement(li)
        })
        fireEvent(links[0],
            new MouseEvent('click', {
                bubbles: true,
                cancelable: true,
            })
        )
        expect(list).not.toBeInTheDocument()
        listItems.forEach((li) => expect(li).not.toBeInTheDocument())
    })
})
