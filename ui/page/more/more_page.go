package more

import (
	"context"

	"gioui.org/layout"

	"github.com/planetdecred/godcr/app"
	"github.com/planetdecred/godcr/ui/decredmaterial"
	"github.com/planetdecred/godcr/ui/load"
	"github.com/planetdecred/godcr/ui/page/components"
	"github.com/planetdecred/godcr/ui/values"
)

const MorePageID = "More"

type (
	C = layout.Context
	D = layout.Dimensions
)

type MorePage struct {
	*load.Load
	// GenericPageModal defines methods such as ID() and OnAttachedToNavigator()
	// that helps this Page satisfy the app.Page interface. It also defines
	// helper methods for accessing the PageNavigator that displayed this page
	// and the root WindowNavigator.
	*app.GenericPageModal

	ctx        context.Context // page context
	ctxCancel  context.CancelFunc
	backButton decredmaterial.IconButton
	infoButton decredmaterial.IconButton
	//container  *widget.List
	shadowBox *decredmaterial.Shadow
}

func NewMorePage(l *load.Load) *MorePage {
	pg := &MorePage{
		Load:             l,
		GenericPageModal: app.NewGenericPageModal(MorePageID),
		shadowBox:        l.Theme.Shadow(),
	}
	pg.backButton, pg.infoButton = components.SubpageHeaderButtons(l)
	pg.backButton.Icon = pg.Theme.Icons.NavigationArrowBack
	return pg
}

// OnNavigatedTo is called when the page is about to be displayed and
// may be used to initialize page features that are only relevant when
// the page is displayed.
// Part of the load.Page interface.

func (pg *MorePage) OnNavigatedTo() {
	pg.ctx, pg.ctxCancel = context.WithCancel(context.TODO())

}

func (pg *MorePage) Layout(gtx layout.Context) layout.Dimensions {
	if pg.Load.GetCurrentAppWidth() <= gtx.Dp(values.StartMobileView) {
		return pg.layoutMobile(gtx)
	}
	return pg.layoutDesktop(gtx)
}

func (pg *MorePage) layoutDesktop(gtx layout.Context) layout.Dimensions {
	container := func(gtx C) D {

		return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return layout.Inset{Bottom: values.MarginPadding16}.Layout(gtx, func(gtx C) D {
					return pg.topNav(gtx)
				})
			}),
			layout.Rigid(func(gtx C) D {
				return layout.Inset{Left: values.MarginPadding50, Bottom: values.MarginPadding16}.Layout(gtx, func(gtx C) D {
					return pg.Theme.Label(values.TextSize16, values.StrMorePageInfo).Layout(gtx)
				})
			}),

			layout.Rigid(func(gtx C) D {

				return layout.Inset{Left: values.MarginPadding100, Right: values.MarginPadding50}.Layout(gtx, func(gtx C) D {
					pg.shadowBox.SetShadowRadius(14)

					return decredmaterial.LinearLayout{
						Width:       decredmaterial.WrapContent,
						Orientation: layout.Vertical,
						Height:      decredmaterial.WrapContent,
						Padding:     layout.UniformInset(values.MarginPadding18),
						Background:  pg.Theme.Color.Surface,
						Alignment:   layout.Middle,
						Shadow:      pg.shadowBox,
						Margin:      layout.UniformInset(values.MarginPadding4),
						Border:      decredmaterial.Border{Radius: decredmaterial.Radius(14)},
					}.Layout(gtx,

						layout.Rigid(func(gtx C) D {
							return layout.Inset{
								Left:  values.MarginPadding150,
								Right: values.MarginPadding150,
								Top:   values.MarginPadding2,
							}.Layout(gtx, pg.Theme.Icons.DocumentationIcon.Layout36dp)

						}),
						layout.Rigid(func(gtx C) D {

							return layout.Inset{
								Left:  values.MarginPadding130,
								Right: values.MarginPadding130,
								Top:   values.MarginPadding2,
							}.Layout(gtx, pg.Theme.Body1(values.String(values.StrDocumentation)).Layout)

						}),
					)

				})

			}),
		)

	}
	return components.UniformPadding(gtx, container)
}

func (pg *MorePage) layoutMobile(gtx layout.Context) layout.Dimensions {
	container := func(gtx C) D {

		return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return layout.Inset{Bottom: values.MarginPadding16}.Layout(gtx, func(gtx C) D {
					return pg.topNav(gtx)
				})
			}),
			layout.Rigid(func(gtx C) D {
				return layout.Inset{Left: values.MarginPadding50, Bottom: values.MarginPadding16}.Layout(gtx, func(gtx C) D {
					return pg.Theme.Label(values.TextSize16, values.StrMorePageInfo).Layout(gtx)
				})
			}),

			layout.Rigid(func(gtx C) D {

				return layout.Inset{Left: values.MarginPadding100, Right: values.MarginPadding50}.Layout(gtx, func(gtx C) D {
					pg.shadowBox.SetShadowRadius(14)

					return decredmaterial.LinearLayout{
						Width:       decredmaterial.WrapContent,
						Orientation: layout.Vertical,
						Height:      decredmaterial.WrapContent,
						Padding:     layout.UniformInset(values.MarginPadding18),
						Background:  pg.Theme.Color.Surface,
						Alignment:   layout.Middle,
						Shadow:      pg.shadowBox,
						Margin:      layout.UniformInset(values.MarginPadding4),
						Border:      decredmaterial.Border{Radius: decredmaterial.Radius(14)},
					}.Layout(gtx,

						layout.Rigid(func(gtx C) D {
							return layout.Inset{
								Left:  values.MarginPadding150,
								Right: values.MarginPadding150,
								Top:   values.MarginPadding2,
							}.Layout(gtx, pg.Theme.Icons.DocumentationIcon.Layout36dp)

						}),
						layout.Rigid(func(gtx C) D {

							return layout.Inset{
								Left:  values.MarginPadding140,
								Right: values.MarginPadding140,
								Top:   values.MarginPadding2,
							}.Layout(gtx, pg.Theme.Body1(values.String(values.StrDocumentation)).Layout)

						}),
					)

				})

			}),
		)

	}
	return components.UniformPadding(gtx, container)
}

// HandleUserInteractions is called just before Layout() to determine
// if any user interaction recently occurred on the page and may be
// used to update the page's UI components shortly before they are
// displayed.
// Part of the load.Page interface.
func (pg *MorePage) HandleUserInteractions() {

}

// OnNavigatedFrom is called when the page is about to be removed from
// the displayed window. This method should ideally be used to disable
// features that are irrelevant when the page is NOT displayed.
// NOTE: The page may be re-displayed on the app's window, in which case
// OnNavigatedTo() will be called again. This method should not destroy UI
// components unless they'll be recreated in the OnNavigatedTo() method.
// Part of the load.Page interface.

func (pg *MorePage) OnNavigatedFrom() {
	pg.ctxCancel()
}

func (pg *MorePage) topNav(gtx C) D {
	m := values.MarginPadding20
	return layout.Flex{}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return pg.backButton.Layout(gtx)
				}),
				layout.Rigid(func(gtx C) D {
					return layout.Inset{Left: m}.Layout(gtx, pg.Theme.H6(values.String(values.StrHelp)).Layout)
				}),
			)
		}),
		layout.Flexed(1, func(gtx C) D {
			return layout.E.Layout(gtx, pg.infoButton.Layout)
		}),
	)

}
