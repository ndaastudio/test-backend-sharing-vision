import * as React from "react";

import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar";
import { Link, useLocation } from "react-router-dom";

const data = {
  // navMain: [
  //   {
  navMain: [
    {
      title: "All Posts",
      url: "/",
    },
    {
      title: "Preview",
      url: "/preview",
    },
  ],
  // },
  // ],
};

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const { pathname } = useLocation();
  const [activeLink, setActiveLink] = React.useState(pathname);

  React.useEffect(() => {
    setActiveLink(pathname);
  }, [pathname]);

  return (
    <Sidebar {...props}>
      <SidebarContent>
        <SidebarGroup>
          {data.navMain.map((item) => (
            <SidebarMenu key={item.title}>
              <SidebarMenuItem>
                <SidebarMenuButton asChild isActive={activeLink === item.url}>
                  <Link to={item.url}>{item.title}</Link>
                </SidebarMenuButton>
              </SidebarMenuItem>
            </SidebarMenu>
          ))}
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  );
}
