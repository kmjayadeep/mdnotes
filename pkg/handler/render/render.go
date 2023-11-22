package render

import(
  	"github.com/gofiber/fiber/v2"
)

type Render struct {
}

func NewRender() *Render {
  return &Render{}
}

func (r *Render) Index(c *fiber.Ctx) error {
  return c.Render("index", fiber.Map{})
}
