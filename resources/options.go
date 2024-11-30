package resources

type Language string

const (
	EnglishUk Language = "en-GB"
)

type Options struct {
	language Language
}

var GameOptions = &Options{
	// TODO: english uk is currently the default language but this should be driven
	// by the user's system settings perhaps.
	language: EnglishUk,
}

func (o *Options) GetLanguage() Language {
	return o.language
}

func (o *Options) SetLanguage(language Language) {
	o.language = language
}
