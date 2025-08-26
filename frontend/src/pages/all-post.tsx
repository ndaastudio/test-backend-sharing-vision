import CreateArticleModal from "@/components/create-article-modal";
import DraftedTable from "@/components/drafted-table";
import PublishedTable from "@/components/published-table";
import TrashedTable from "@/components/trashed-table";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

export default function AllPost() {
  return (
    <>
      <CreateArticleModal />

      <Tabs defaultValue="published">
        <div className="flex justify-center">
          <TabsList>
            <TabsTrigger value="published">Published</TabsTrigger>
            <TabsTrigger value="draft">Draft</TabsTrigger>
            <TabsTrigger value="trashed">Trashed</TabsTrigger>
          </TabsList>
        </div>

        <TabsContent value="published">
          <PublishedTable />
        </TabsContent>
        <TabsContent value="draft">
          <DraftedTable />
        </TabsContent>
        <TabsContent value="trashed">
          <TrashedTable />
        </TabsContent>
      </Tabs>
    </>
  );
}
