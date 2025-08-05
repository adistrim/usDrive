import { Button } from '@/components/ui/button'
import { createFileRoute } from '@tanstack/react-router'
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"

export const Route = createFileRoute('/')({
    component: RouteComponent,
})

function RouteComponent() {
    return (
    <div className="flex justify-end p-5 top-0 items-center pr-16">
      <Dialog>
        <form>
          <DialogTrigger asChild>
            <Button className="cursor-pointer ">Upload files</Button>
          </DialogTrigger>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>Upload your files here!</DialogTitle>
            <div>
                <Label htmlFor="file-upload">File</Label>
                <Input id="file-upload" type="file" />
            </div>
          </DialogHeader>
          <DialogFooter>
            <DialogClose asChild>
              <Button type="submit">Upload</Button>
            </DialogClose>
          </DialogFooter>
        </DialogContent>
      </form>
    </Dialog>
    </div>
  )
}
