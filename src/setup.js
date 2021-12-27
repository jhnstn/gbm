/**
 * WordPress dependencies
 */
import { addAction, addFilter } from '@wordpress/hooks';
import { initialHtmlGutenberg } from '@wordpress/react-native-editor';

/**
 * Internal dependencies
 */
import correctTextFontWeight from './text-font-weight-correct';
import {
	registerJetpackBlocks,
	registerJetpackEmbedVariations,
	setupJetpackEditor,
} from './jetpack-editor-setup';
import { setupBlockExperiments } from './block-experiments-setup';
import initialHtml from './initial-html';
import initAnalytics from './analytics';

const setupHooks = () => {
	// Hook triggered before the editor is rendered
	addAction( 'native.pre-render', 'gutenberg-mobile', ( props ) => {
		const { jetpackState } = props;

		require( './strings-overrides' );
		correctTextFontWeight();

		setupJetpackEditor(
			jetpackState || { blogId: 1, isJetpackActive: true }
		);

		// Jetpack Embed variations use WP hooks that are attached to
		// block type registration, so it’s required to add them before
		// the core blocks are registered.
		registerJetpackEmbedVariations( props );
	} );

	addFilter(
		'native.block_editor_props',
		'gutenberg-mobile',
		( editorProps ) => {
			if ( __DEV__ ) {
				let { initialTitle, initialData } = editorProps;

				if ( initialTitle === undefined ) {
					initialTitle = 'Welcome to gutenberg for WP Apps!';
				}

				if ( initialData === undefined ) {
					initialData = initialHtml + initialHtmlGutenberg;
				}

				return {
					...editorProps,
					initialTitle,
					initialData,
				};
			}
			return editorProps;
		}
	);

	// Hook triggered after the editor is rendered
	addAction( 'native.render', 'gutenberg-mobile', ( props ) => {
		const capabilities = props.capabilities ?? {};
		registerJetpackBlocks( props );
		setupBlockExperiments( capabilities );
	} );
};

export default () => {
	initAnalytics();
	setupHooks();
};
