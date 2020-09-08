package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"pahamify/project/request"
	"pahamify/project/response"
	"pahamify/project/service"
)

type pokemonController struct {
	pokemonService service.PokemonService
}

// PokemonController interface
type PokemonController interface {
	// CreatePokemon insert a new pokemon data
	CreatePokemon(c echo.Context) error
	// UpdatePokemon update an existing pokemon data
	UpdatePokemon(c echo.Context) error
	// GetPokemons retrieve pokemon data
	GetPokemons(c echo.Context) error
	// DeletePokemon remove an existing pokemon data
	DeletePokemon(c echo.Context) error
}

// NewPokemonController routes endpoints related to pokemon requests
func NewPokemonController(
	e *echo.Echo,
	svc *service.Container,
) PokemonController {
	ctrl := &pokemonController{
		pokemonService: svc.PokemonService,
	}

	v1Route := e.Group("/v1")
	v1PokemonRoute := v1Route.Group("/pokemon")
	{
		v1PokemonRoute.POST("", ctrl.CreatePokemon)
		v1PokemonRoute.PUT("", ctrl.UpdatePokemon)
		v1PokemonRoute.GET("", ctrl.GetPokemons)
		v1PokemonRoute.DELETE("", ctrl.DeletePokemon)
	}

	return ctrl
}

func (ctrl *pokemonController) CreatePokemon(c echo.Context) error {
	var req request.Pokemon
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "invalid pokemon request",
		})
	}

	pokemon := req.ToPokemonModel()
	if err := ctrl.pokemonService.CreatePokemon(c.Request().Context(), &pokemon); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("success create pokemon %s", pokemon.Name),
	})
}

func (ctrl *pokemonController) UpdatePokemon(c echo.Context) error {
	var req request.Pokemon
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "invalid pokemon request",
		})
	}

	if req.ID == "" {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "pokemon ID is required",
		})
	}

	pokemon := req.ToPokemonModel()
	if err := ctrl.pokemonService.UpdatePokemon(c.Request().Context(), &pokemon); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("success update pokemon %s", pokemon.Name),
	})
}

func (ctrl *pokemonController) GetPokemons(c echo.Context) error {
	limitQuery := c.QueryParam("limit")
	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("invalid limit %s", limitQuery),
		})
	}

	pokemons, err := ctrl.pokemonService.GetPokemons(c.Request().Context(), limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	if len(pokemons) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "no pokemon data found",
		})
	}

	resp := response.FromPokemonsModel(pokemons)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success retrieve pokemons",
		"data":    resp,
	})
}

func (ctrl *pokemonController) DeletePokemon(c echo.Context) error {
	var req request.Pokemon
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "invalid pokemon request",
		})
	}

	if req.ID == "" {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "pokemon ID is required",
		})
	}

	pokemon := req.ToPokemonModel()
	if err := ctrl.pokemonService.DeletePokemon(c.Request().Context(), &pokemon); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("success delete pokemon %s", pokemon.Name),
	})
}
