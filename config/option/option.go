package option

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Key string

var (
	ErrInvalidOption = errors.New("invalid option")
	ErrKeyNotFound   = errors.New("key not found")
	ErrInvalidKey    = errors.New("invalid key")
)

const NoConfigMessage = "not configured"

var ErrNotDirectory = errors.New("not a directory")

type Option struct {
	key       string
	viperKey  string
	getString func() string
	validate  func(string) error
}

func (o *Option) Key() string {
	return o.key
}

func (o *Option) String() string {
	return o.getString()
}

func (o *Option) Validate(v string) error {
	if o.validate == nil {
		return nil
	}

	return o.validate(v)
}

func (o *Option) SetValue(v string) error {
	viper.Set(o.viperKey, v)
	return viper.WriteConfig()
}

type Builder struct {
	option Option
}

func NewOptionBuilder(key string) *Builder {
	return &Builder{
		option: Option{
			key:      key,
			viperKey: key,
		},
	}
}

func (b *Builder) WithViperKey(v string) *Builder {
	b.option.viperKey = v
	return b
}

func (b *Builder) WithGetString(f func() string) *Builder {
	b.option.getString = f
	return b
}

func (b *Builder) WithValidate(f func(string) error) *Builder {
	b.option.validate = f
	return b
}

func (b *Builder) Build() Option {
	if b.option.getString == nil {
		b.option.getString = func() string {
			v := viper.GetString(b.option.viperKey)
			if v == "" {
				return NoConfigMessage
			}

			return v
		}
	}

	return b.option
}

type Registry struct {
	options map[string]Option
}

func NewRegistry() *Registry {
	registry := &Registry{
		options: make(map[string]Option),
	}

	registry.registerAll()

	return registry
}

func (r *Registry) registerAll() {
	editorOption := NewOptionBuilder("editor").Build()
	r.register(editorOption)

	workspaceOption := NewOptionBuilder("workspace").
		WithValidate(func(s string) error {
			if f, err := os.Stat(s); os.IsNotExist(err) || !f.IsDir() {
				return ErrNotDirectory
			}

			return nil
		}).Build()
	r.register(workspaceOption)
}

func (r *Registry) register(option Option) {
	r.options[option.Key()] = option
}

func (r *Registry) Get(key string) (Option, error) {
	option, ok := r.options[key]
	if !ok {
		return Option{}, fmt.Errorf("%w: %w: %s", ErrInvalidOption, ErrKeyNotFound, key)
	}

	return option, nil
}

func (r *Registry) GetAll() map[string]Option {
	return r.options
}
